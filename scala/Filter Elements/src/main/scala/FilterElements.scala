object FilterElements {
  def solve(arr:Array[Int], atLeast:Int):Unit={
    val res = reduce(arr, atLeast)
    res.length match {
      case 0 => println("-1")
      case _ => println(res.mkString(" "))
    }
  }

  def reduce(arr:Array[Int], atLeast:Int):Array[Int]= for(de <- arr.distinct if arr.count(_==de)>=atLeast) yield de

  def main(args:Array[String]):Unit={
    val nTestCases = scala.io.StdIn.readInt()
    var atLeast :Int = 0
    var arr :Array[Int] = Array()
    for(i<-0 until nTestCases){
      atLeast = scala.io.StdIn.readLine().split(" ")(1).toInt
      arr = scala.io.StdIn.readLine().split(" ").map(_.toInt)
      solve(arr, atLeast)
    }

    //val arr = scala.io.StdIn.readLine().split(" ").map(_.toInt)
    //reduce(Array(1,1,2,3,3,4,5,6,6,7,7),2).foreach(x => println(x))
    //solve(Array(5,4,3,2,1,1,2,3,4,5),2)
  }

}
