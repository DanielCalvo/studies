
import groovy.json.JsonSlurper

def url = new URL("http://a.4cdn.org/o/threads.json" )
def connection = url.openConnection()

threads = connection.content

def ui = new JsonSlurper().parse(threads)

ui.items.each

// get the JSON response
//def json = connection.inputStream.withCloseable { inStream -> new JsonSlurper().parse( inStream as InputStream ) }