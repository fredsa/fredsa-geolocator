import webapp2

X_APPENGINE_COUNTRY = 'X-AppEngine-Country'
X_APPENGINE_REGION = 'X-AppEngine-Region'
X_APPENGINE_CITY = 'X-AppEngine-City'
X_APPENGINE_CITYLATLONG = 'X-AppEngine-CityLatLong'

class MainHandler(webapp2.RequestHandler):

  def get(self):
    location = self.request.headers.get(X_APPENGINE_CITYLATLONG)
    self.response.write(location)

app = webapp2.WSGIApplication([
  ('/.*', MainHandler),
], debug=True)
