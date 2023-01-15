import React, { useState } from "react";
import SharedHeader from "../shared/SharedHeader";
import { useNewPost } from "../hooks/useNewPost";
import { useNavigate } from "react-router-dom";

function Talk() {
  const postNewPost = useNewPost();
  const navigate = useNavigate();

  const [input, setInput] = useState("");
  const sleep = (ms) => new Promise((r) => setTimeout(r, ms));
  const [loading, setLoading] = useState(false);

  return (
    <div>
      <SharedHeader />
      <div style={{ display: "flex", flexDirection: "column" }}>
        <div
          style={{
            marginTop: 0,
            color: "#C4BAA6",
            fontFamily: "theseasons",
            alignSelf: "center",
            justifyContent: "center",
            textAlign: "center",
            fontSize: 63,
          }}
        >
          How are <span style={{ fontStyle: "italic" }}>you</span> feeling
          today?
        </div>

        <textarea
          placeholder="Type something here..."
          style={{
            color: "#C4BAA6",
            fontFamily: "theseasons",
            background: "transparent",
            alignSelf: "center",
            justifyContent: "center",
            textAlign: "center",
            fontSize: 24,
            border: "1px solid #C4BAA6",
            width: 600,
            height: 350,
            padding: "10px 60px 10px 60px",
            borderRadius: "20px",
            cursor: "pointer",
            overflow: "scroll",
            resize: "none",
          }}
          value={input}
          onChange={(e) => setInput(e.target.value)}
        />
        <div
          style={{
            color: "#C4BAA6",
            fontFamily: "theseasons",
            background: "transparent",
            alignSelf: "center",
            textAlign: "center",
            fontSize: 24,
            width: 300,
            height: 300,
            cursor: "pointer",
          }}
          onClick={async () => {
            console.log("input is ", input);
            postNewPost.mutate({
              body: input,
            });
            setLoading(true);
            await sleep(1000);
            navigate("/feed");
          }}
        >
          {loading === true ? "Loading..." : "Submit"}
        </div>
      </div>
    </div>
  );
}

export default Talk;
