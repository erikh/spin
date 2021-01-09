import React from "react";
import "./App.css";
import RFB from "@novnc/novnc";
import sprintf from 'sprintf';

class App extends React.Component {
  componentDidMount() {
    var id = parseInt(new URLSearchParams(window.location.search).get("id"));

    var rfb = new RFB(
      document.getElementById("root"),
      sprintf("ws://127.0.0.1:60%02d", id)
    );
    rfb.resizeSession = true;
    rfb.scaleViewport = true;
  }
  render() {
    return <div />;
  }
}

export default App;
