object PythagoreanTriplet {
  def isPythagorean(t :(Int,Int,Int)):Boolean ={   
    val a = t._1
    val b = t._2
    val c = t._3
    ( (a*a) + (b*b) == c*c )
  }
  def pythagoreanTriplets(f :Int, t :Int):Seq[(Int,Int,Int)] = {
    for{
      i <- f to t
      j <- i+1 to t
      k <- j+1 to t
      if isPythagorean(i,j,k)
    }yield (i,j,k)
  }

  def pythagoreanTripletsSum(sum:Int):Seq[(Int,Int,Int)] = {
    for{
      i <- 1 to sum/3
      j <- i to (sum-i)/2
      if isPythagorean(i,j,sum-(i+j))
    } yield (i,j,sum-i-j)
  }

  def main(args:Array[String]):Unit={
    pythagoreanTriplets(0,1000).foreach(println)
    pythagoreanTripletsSum(180).foreach(println)
  }

}
