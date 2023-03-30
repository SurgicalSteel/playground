object Bob {
  def response(statement: String): String = statement.trim() match {
    case statement if statement.isEmpty => "Fine. Be that way!"
    case statement if isWholeStringUppercase(statement) => if (statement.endsWith("?"))  "Calm down, I know what I'm doing!" else "Whoa, chill out!"
    case statement if statement.endsWith("?") => "Sure."
    case _ => "Whatever."
  }

  def isWholeStringUppercase(str:String):Boolean = str.exists(subStr => subStr.isLetter) &&  str.toUpperCase() == str

  def main(args:Array[String]):Unit={
    println(response("HA?"))
  }
}
