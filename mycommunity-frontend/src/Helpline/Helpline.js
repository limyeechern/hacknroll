import React from "react";
import Header from "../Home/Header";
import { useNavigate } from "react-router-dom";
import HomeOutlinedIcon from "@mui/icons-material/HomeOutlined";

function Helpline() {
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
        It's <span style={{ fontStyle: "italic" }}>okay</span> to reach out.
      </div>
      <div
        style={{
          color: "#C4BAA6",
          // fontFamily: "theseasons",
          alignSelf: "center",
          textAlign: "center",
          fontSize: 20,
          position: "absolute",
          top: "300px",
          padding: "50px",
          width: "900px",
        }}
      >
        National Care Hotline: 1800-202-6868
        <br />
        <br />
        Mental Well-being - Institute of Mental Health’s Mental Health Helpline
        (6389-2222) - Samaritans of Singapore (1800-221-4444) - Silver Ribbon
        Singapore (6385-3714)
        <br />
        <br />
        Counselling - TOUCHline (Counselling) – 1800 377 2252
        <br />
        <br /> For other Helplines and mental health-related support -
        go.gov.sg/hotlines (BELLE, Beyond the Label helpbot) -
        www.msf.gov.sg/Pages/Contact-Us.aspx
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

export default Helpline;
