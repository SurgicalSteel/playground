object BeerSong{
  def getPlural(count :Int, isCapital:Boolean=false):String = count match{
    case count if count>1 => s"$count bottles"
    case count if count==1 => s"$count bottle"
    case count if count ==0 && isCapital => "No more bottles"
    case _ => "no more bottles"
  }

  def getFirst(count:Int):String = count match {
    case 0 => s"${getPlural(count, true)} of beer on the wall, ${getPlural(count)} of beer.\n"
    case _ => s"${getPlural(count)} of beer on the wall, ${getPlural(count)} of beer.\n"
  }

  def getSecond(count:Int):String = {
    var nextCount = count -1
    count match {
      case 1 => s"Take it down and pass it around, ${getPlural(nextCount)} of beer on the wall.\n"
      case 0 => s"Go to the store and buy some more, ${getPlural(99)} of beer on the wall.\n"
      case _ => s"Take one down and pass it around, ${getPlural(nextCount)} of beer on the wall.\n"
    }
  }

  def recite(bottleCount:Int, reciteCount:Int ):String={
    var currentCount :Int = bottleCount
    var iterator :Int =0
    var result :String = ""
    for (iterator <- 0 until reciteCount){
      result = result + getFirst(currentCount)+getSecond(currentCount)
      if (iterator +1 < reciteCount){
          result += "\n"
      }

      currentCount -= 1
    }
    result
  }

  def main(args:Array[String]):Unit={
    println(recite(99,2))
  }

}