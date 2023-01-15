import React from "react";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import Typography from "@mui/material/Typography";
import Modal from "@mui/material/Modal";
import { Link } from "react-router-dom";
import CloseIcon from "@mui/icons-material/Close";

const style = {
  position: "absolute",
  color: "#F4F1ED",
  top: "50%",
  left: "50%",
  transform: "translate(-50%, -50%)",
  minHeight: 300,
  width: 600,
  bgcolor: "rgba(20,20,20,20.8)",
  border: "2px solid #000",
  boxShadow: 24,
  p: 4,
  borderRadius: "30px",
  display: "flex",
  flexDirection: "column",
};

function Tag({ data }) {
  const [open, setOpen] = React.useState(false);
  const handleOpen = () => setOpen(true);
  const handleClose = () => setOpen(false);
  console.log("data is ", data);
  return (
    <>
      <div
        style={{
          border: "1px solid #C4BAA6",
          fontSize: "20px",
          padding: "0px 7px 0px 7px",
          marginRight: "5px",
          borderRadius: "10px",
          backgroundColor: "rgb(50,50,50,0.5)",
          cursor: "pointer",
        }}
        onClick={handleOpen}
      >
        {data.title}
      </div>
      <Modal
        open={open}
        onClose={handleClose}
        aria-labelledby="modal-modal-title"
        aria-describedby="modal-modal-description"
      >
        <Box sx={style}>
          <CloseIcon
            onClick={handleClose}
            style={{
              cursor: "pointer",
              alignSelf: "flex-end",
              marginTop: "-10px",
            }}
          />

          <div
            style={{
              display: "flex",
              flexDirection: "row",
              justifyContent: "space-between",
            }}
          >
            <div style={{ display: "flex", flexDirection: "column" }}>
              <Typography id="modal-modal-title" variant="h6" component="h2">
                {data.title}
              </Typography>
              <div
                onClick={() => {
                  window.open(data.url, "_blank");
                }}
                style={{ cursor: "pointer" }}
              >
                <Typography
                  id="modal-modal-description"
                  style={{ fontSize: "12px" }}
                >
                  Go to subreddit
                </Typography>
              </div>
            </div>
            <Typography
              id="modal-modal-description"
              style={{ fontSize: "12px" }}
            >
              {`${data.subscribers} members`}
            </Typography>
          </div>
          <Typography id="modal-modal-description" sx={{ mt: 2 }}>
            {data.public_description}
          </Typography>
        </Box>
      </Modal>
    </>
  );
}

export default Tag;
