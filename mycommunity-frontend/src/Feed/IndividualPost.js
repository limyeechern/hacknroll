import { display } from "@mui/system";
import React from "react";
import Tag from "./Tag";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import Typography from "@mui/material/Typography";
import Modal from "@mui/material/Modal";

function IndividualPost({ data }) {
  console.log(data);
  return (
    <div
      style={{
        width: "600px",
        border: "1px solid rgba(255,255,255,0.3)",
        // backgroundColor: "rgba(255,255,255,0.5)",
        borderRadius: "20px",
        minHeight: "200px",
        alignSelf: "center",
        marginTop: "20px",
        color: "#C4BAA6",
        fontFamily: "cocogothic",
        textAlign: "center",
        fontSize: 20,
        margin: "20px",
        display: "inline-block",
        flexDirection: "column",
        overflow: "scroll",
      }}
    >
      <div
        style={{
          padding: " 10px",
          alignSelf: "flex-start",
          display: "flex",
          flexDirection: "row",
          justifyContent: "flex-start",
          width: "95%",
          overflow: "auto",
          "::-webkit-scrollbar": { display: "none" },
        }}
      >
        <Tag data={data.reddit_data[0]} />
        <Tag data={data.reddit_data[1]} />
        <Tag data={data.reddit_data[2]} />
        <Tag data={data.reddit_data[3]} />
        <Tag data={data.reddit_data[4]} />
      </div>
      <div style={{ color: "white", padding: "20px" }}>{data.body}</div>
    </div>
  );
}

export default IndividualPost;
