object Fibonacci {

  def fibonacci(n:Int, m :Map[Int,BigInt]):BigInt = m(n-1) + m(n-2)

  def main(args:Array[String]):Unit={
    var m :Map[Int,BigInt] = Map()
    m = m updated(0,0)
    m = m updated(1,1)

    for(i <- 2 to 10000){
      m = m updated(i, fibonacci(i, m))
    }

    val nTestCases = scala.io.StdIn.readInt()
    var n :Int=0

    for(i<- 0 until nTestCases){
      n = scala.io.StdIn.readInt()
      println(m(n)%100000007)
    }

  }

}
