class App extends React.Component {
  render() {
    if (this.loggedIn) {
      return (<LoggedIn />);
    } else {
      return (<Home />);
    }
  }
}

class Home extends React.Component {
  render() {
    return (
      <div className="container">
        <div className="col-xs-8 col-xs-offset-2 text-center">
          <h1> Hey There! :) </h1>
        </div>
      </div>
    )
  }
}

ReactDOM.render(<App />, document.getElementById('app'));
