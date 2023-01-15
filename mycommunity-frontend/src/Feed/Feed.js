import React, { useState } from "react";
import SharedHeader from "../shared/SharedHeader";
import { useFeed } from "../hooks/useFetchPost";
import IndividualPost from "./IndividualPost";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";

function Feed() {
  const {
    data,
    isSuccess,
    fetchNextPage,
    hasNextPage,
    isFetchingNextPage,
    isLoading,
  } = useFeed();

  const [loading, setLoading] = useState(false);

  const sleep = (ms) => new Promise((r) => setTimeout(r, ms));

  return (
    <div style={{ height: "100vh" }}>
      <SharedHeader />
      <div style={{ display: "flex", flexDirection: "column" }}>
        <div
          style={{
            marginTop: -100,
            color: "#C4BAA6",
            fontFamily: "theseasons",
            textAlign: "center",
            fontSize: 63,
            width: "50vw",
            alignSelf: "center",
          }}
        >
          Hey, you are not <span style={{ fontStyle: "italic" }}>alone.</span>
        </div>
        {isLoading ? (
          <div style={{ color: "white" }}>Loading</div>
        ) : (
          <>
            <div
              style={{
                height: "80vh",
                display: "flex",
                flexDirection: "column",
                overflow: "scroll",
              }}
            >
              {data.pages.map((page, i) => (
                <React.Fragment key={i}>
                  {page.data.map((item) =>
                    item.reddit_data !== null ? (
                      <IndividualPost key={item._id} data={item} />
                    ) : null
                  )}
                </React.Fragment>
              ))}
            </div>
            <div
              style={{
                color: "#C4BAA6",
                fontFamily: "theseasons",
                textAlign: "center",
                fontSize: 20,
                cursor: "pointer",
                display: "flex",
                alignSelf: "center",
              }}
              onClick={async () => {
                setLoading(true);
                await sleep(500);
                fetchNextPage();
                setLoading(false);
              }}
            >
              {loading ? "Loading..." : "Click for more"}
              <ExpandMoreIcon />
            </div>
          </>
        )}
      </div>
    </div>
  );
}

export default Feed;
