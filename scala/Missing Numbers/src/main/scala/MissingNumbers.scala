

object MissingNumbers {

  def solve(arra :Array[Int], arrb :Array[Int]):String={
    var mb :Map[Int,Int] = arrb.groupBy(l => l).map(t => (t._1,arrb.count(_ == t._1)))
    var ma :Map[Int,Int] = arra.groupBy(l => l).map(t => (t._1,arra.count(_ == t._1)))

    var res :Array[Int] = Array()
    for(i <- mb.keys){
      if (!ma.contains(i)){
        for(j<- 0 until mb(i)){
          res = res :+ i
        }
      }else{
        if(ma(i)!=mb(i)){
          for(j<- 0 until mb(i) - ma(i)){
            res = res :+ i
          }
        }
      }
    }
    res.distinct.sorted.mkString(" ")
  }


  def main(args:Array[String]):Unit={
    val sza = scala.io.StdIn.readLine().toInt
    val arra = scala.io.StdIn.readLine().split(" ").map(_.toInt)
    val szb = scala.io.StdIn.readLine().toInt
    val arrb = scala.io.StdIn.readLine().split(" ").map(_.toInt)
    println(solve(arra,arrb))
  }

}
