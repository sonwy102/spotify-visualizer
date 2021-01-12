function App() {
  return (
    <Router>
      <Container fluid className="App">
          <h1>Twitter-Spotify Curator</h1>
      </Container>

      <div>
        <Switch>
          <Route path="/" exact={true}>
            <Home />
          </Route>
        </Switch>
      </div>
    </Router>
  )
}
