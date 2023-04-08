object SuperDigit {

  def getSumDigit(s :String):Long = {
    var sum :Long = 0
    s.foreach(c => sum += c.toString.toInt)
    sum
  }

  def main(args:Array[String]):Unit={
    val s = scala.io.StdIn.readLine()
    val n :String = s.split(" ")(0)
    val k :Int = s.split(" ")(1).toInt
    var res:Long = getSumDigit(n) * k
    if(res <10){
      println(res)

    }else{
      while(res>=10){
        res = getSumDigit(res.toString)
      }
      println(res)
    }

  }
}
