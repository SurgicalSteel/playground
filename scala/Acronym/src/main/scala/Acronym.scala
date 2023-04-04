object Acronym {
  def abbreviate(phrase:String):String={
    var res = ""
    phrase.split("\\s+|\\-").filterNot(_.isEmpty).foreach(x => res += x.charAt(0).toString.capitalize)
    res
  }
  def main(args :Array[String]):Unit={
    println(abbreviate("Something - I made up from thin air"))
  }
}
