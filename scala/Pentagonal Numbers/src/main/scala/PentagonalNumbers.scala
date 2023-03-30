object PentagonalNumbers {
  def main(args: Array[String]) {
    var m :Map[Int,Long] = Map()
    m += (1 -> 1)
    m+= (2 -> 5)
    var currAdd :Long = 7
    for(i<-3 to 500000){
      m += (i -> (m(i-1) + currAdd))
      currAdd += 3
    }
    var n :Int = scala.io.StdIn.readInt()
    var temp :Int = 0
    for(i<-1 to n){
      temp = scala.io.StdIn.readInt()
      println(m(temp))
    }
  }
}
