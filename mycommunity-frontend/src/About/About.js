import React from "react";
import Header from "../Home/Header";
import { useNavigate } from "react-router-dom";
import HomeOutlinedIcon from "@mui/icons-material/HomeOutlined";

function About() {
  let navigate = useNavigate();

  return (
    <div style={{ display: "flex", flexDirection: "column" }}>
      <Header />
      <div
        style={{
          marginTop: 0,
          color: "#C4BAA6",
          fontFamily: "theseasonslight",
          alignSelf: "center",
          justifyContent: "center",
          textAlign: "center",
          fontSize: 90,
          position: "absolute",
          top: "180px",
        }}
      >
        About us
      </div>
      <div
        style={{
          color: "#C4BAA6",
          // fontFamily: "theseasons",
          alignSelf: "center",
          textAlign: "center",
          fontSize: 30,
          position: "absolute",
          top: "300px",
          padding: "50px",
          width: "900px",
        }}
      >
        Sometimes, all we need is a listening ear, a familiar voice and a
        community that understands us.
        <br />
        <br />
        MyCommunity strives to provide you with the freedom, space and privacy
        to express yourself and find a safe community you can reside in.
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
        onClick={() => navigate("/")}
      >
        Go Home
      </div>
    </div>
  );
}

export default About;
