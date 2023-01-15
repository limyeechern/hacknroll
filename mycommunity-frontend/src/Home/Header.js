import React from "react";
import { useNavigate } from "react-router-dom";

function Header() {
  const navigate = useNavigate();
  return (
    <div
      style={{
        color: "#C4BAA6",
        display: "flex",
        flexDirection: "row",
        justifyContent: "space-around",
        paddingTop: "50px",
      }}
    >
      <div
        style={{
          border: "1px solid #C4BAA6",
          borderRadius: "30px",
          padding: "10px 80px 10px 80px",
          cursor: "pointer",
        }}
        onClick={() => navigate("/helpline")}
      >
        HELP LINE
      </div>
      <div
        style={{
          border: "1px solid #C4BAA6",
          borderRadius: "30px",
          padding: "10px 80px 10px 80px",
          cursor: "pointer",
        }}
        onClick={() => navigate("/feed")}
      >
        FIND STORIES
      </div>
      <div
        style={{
          border: "1px solid #C4BAA6",
          borderRadius: "30px",
          padding: "10px 80px 10px 80px",
          cursor: "pointer",
        }}
        onClick={() => {
          navigate("/about");
        }}
      >
        ABOUT US
      </div>
    </div>
  );
}

export default Header;
