import React from "react";
import "./App.css";
import RFB from "@novnc/novnc";

class App extends React.Component {
  componentDidMount() {
    var id = new URLSearchParams(window.location.search).get("id");

    var rfb = new RFB(
      document.getElementById("root"),
      `ws://127.0.0.1:60${id}`
    );
    rfb.resizeSession = true;
    rfb.scaleViewport = true;
  }
  render() {
    return <div />;
  }
}

export default App;
