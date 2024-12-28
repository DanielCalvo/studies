#rdoc generates documentation for you based on ruby comments!
#This will generate documentation for the Album class on rdoc. Run `rdoc album.rb`
class Album
  include Enumerable

  #An array of songs. Each song should be a string
  attr_reader :songs

  #Instantiates an Album object with no starter songs
  def initialize
    @songs = []
  end

  #Add a song to the album's song collection
  def add_song(song)
    songs << song
  end

  #Iterate over each song on the album.
  def each
    songs.each { |song| yield song }
  end
end