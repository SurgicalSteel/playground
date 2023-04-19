import sttp.client4._
import net.liftweb.json._

case class Genre(id:Long, name:String)
case class GenreCollection(genres:Array[Genre])
object tmdb_playground {
  def getBearerToken(token :String):String = s"Bearer $token"

  def main(args:Array[String]):Unit={
    val bearerAuthToken = getBearerToken(System.getenv("TMDB_TOKEN_V4"))
    val request = basicRequest.get(uri"https://api.themoviedb.org/3/genre/movie/list").header("Authorization",bearerAuthToken)

    val backend = DefaultSyncBackend()
    val response = request.send(backend)
    implicit val formats = net.liftweb.json.DefaultFormats

    response.body match {
      case Left(s) => println(s"Error: $s")
      case Right(i) => {
        val rawJson = parse(i)
        val genreCollection = rawJson.extract[GenreCollection]
        for(genre <- genreCollection.genres){println(s"${genre.id} ${genre.name}")}
      }
    }

  }
}
