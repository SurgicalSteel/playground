object NBST {
  var m :Map[BigInt, BigInt] = Map()

  def calculate(n:BigInt):BigInt ={
    if(n<=1){
      return 1
    }
    val pred :BigInt = n-1
    var sum :BigInt  = 0
    var left :BigInt = 0
    var right :BigInt = 0
    for(i<- 0 to pred.toInt){
      if(m.contains(i)){
        left = m(i)
      }else{
        left = calculate(i)
        m = m updated(i,left)
      }
      if(m.contains(pred-i)){
        right = m(pred-i)
      }else{
        right = calculate(pred-i)
        m = m updated(pred-i, right)
      }
      sum += (left * right)
    }
    return sum
  }

  def main(args:Array[String]):Unit={
    val nTestCases :Int = scala.io.StdIn.readInt()
    var temp :Int = 0
    for(i<-0 until nTestCases){
      temp= scala.io.StdIn.readInt()
      println(calculate(temp)%100000007)
    }
    //println(calculate(100)%100000007)
  }
}
