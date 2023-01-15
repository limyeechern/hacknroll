import React from "react";
import HomeOutlinedIcon from "@mui/icons-material/HomeOutlined";
import InfoOutlinedIcon from "@mui/icons-material/InfoOutlined";
import { useNavigate } from "react-router-dom";

function SharedHeader() {
  const navigate = useNavigate();
  return (
    <div
      style={{
        color: "white",
        display: "flex",
        justifyContent: "space-between",
        padding: "50px",
      }}
    >
      <div onClick={() => navigate("/")}>
        <HomeOutlinedIcon
          style={{
            color: "#C4BAA6",
            fontSize: "60px",
            cursor: "pointer",
            zIndex: "1",
          }}
        />
      </div>
      <div onClick={() => navigate("/about")}>
        <InfoOutlinedIcon
          style={{ color: "#C4BAA6", fontSize: "60px", cursor: "pointer" }}
        />
      </div>
    </div>
  );
}

export default SharedHeader;
