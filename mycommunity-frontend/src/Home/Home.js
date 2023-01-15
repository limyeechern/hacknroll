import React from "react";
import Header from "./Header";
import { useNavigate } from "react-router-dom";

function Home() {
  let navigate = useNavigate();

  return (
    <div style={{ display: "flex", flexDirection: "column" }}>
      <Header />
      <div
        style={{
          marginTop: 0,
          color: "#C4BAA6",
          fontFamily: "theseasons",
          alignSelf: "center",
          justifyContent: "center",
          textAlign: "center",
          fontSize: 180,
          position: "absolute",
          top: "270px",
        }}
      >
        Hello.
      </div>
      <div
        style={{
          color: "#C4BAA6",
          fontFamily: "theseasons",
          alignSelf: "center",
          textAlign: "center",
          fontSize: 40,
          position: "absolute",
          top: "450px",
        }}
      >
        Welcome to MyCommunity
      </div>
      <div
        style={{
          color: "#C4BAA6",
          alignSelf: "center",
          justifyContent: "center",
          textAlign: "center",
          fontSize: 16,
          border: "1px solid #C4BAA6",
          width: 300,
          padding: "10px 60px 10px 60px",
          borderRadius: "30px",
          cursor: "pointer",
          position: "absolute",
          top: "700px",
        }}
        onClick={() => navigate("/talk")}
      >
        Talk to us
      </div>
    </div>
  );
}

export default Home;
