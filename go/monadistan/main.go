package main

import "fmt"

type ResultWithLog struct {
	result int64
	logs   []string
}

func (rwl *ResultWithLog) GetResult() int64 {
	return rwl.result
}

func (rwl *ResultWithLog) GetLogs() []string {
	return rwl.logs
}

func (rwl *ResultWithLog) PrintResult() {
	fmt.Println(rwl.result)
}

func (rwl *ResultWithLog) PrintLogs() {
	for _, l := range rwl.logs {
		fmt.Println(l)
	}
}

func Square(result int64) ResultWithLog {
	newResult := result * result
	newLogs := []string{
		fmt.Sprintf("number %d is being squared to get %d", result, newResult),
	}

	return ResultWithLog{result: newResult, logs: newLogs}
}

func AddOne(result int64) ResultWithLog {
	newResult := result + 1
	newLogs := []string{
		fmt.Sprintf("number %d is being added by 1 to get %d", result, newResult),
	}

	return ResultWithLog{result: newResult, logs: newLogs}
}

func WrapWithLog(result int64) ResultWithLog {
	return ResultWithLog{
		result: result,
		logs:   make([]string, 0),
	}
}

type Transformer func(result int64) ResultWithLog

func RunSingleTransformation(rwl ResultWithLog, transformerFunc Transformer) ResultWithLog {
	transformedRWL := transformerFunc(rwl.GetResult())

	return ResultWithLog{
		result: transformedRWL.GetResult(),
		logs:   append(rwl.logs, transformedRWL.GetLogs()...),
	}
}

func RunMultipleTransformations(rwl ResultWithLog, transformerFuncs []Transformer) ResultWithLog {
	for _, transformerFunc := range transformerFuncs {
		rwl = RunSingleTransformation(rwl, transformerFunc)
	}
	return rwl
}

func main() {
	rwl := WrapWithLog(2)

	rwl = RunMultipleTransformations(rwl, []Transformer{AddOne, Square, AddOne, Square})

	rwl.PrintResult()
	rwl.PrintLogs()
}
