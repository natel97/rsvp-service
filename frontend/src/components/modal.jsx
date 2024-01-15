import { useEffect, useState } from "react";
import {
  ActionButton,
  ModalBackground,
  ModalBody,
  PageHeader,
} from "../components";

const NotificationListener = () => {
  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");

  const close = () => {
    setDescription("");
    setTitle("");
  };

  useEffect(() => {
    if ("serviceWorker" in navigator) {
      navigator.serviceWorker.addEventListener("message", (message) => {
        console.log({ message });
        if (message.data.type === "OPEN_MODAL") {
          setDescription(message.data.data.body);
          setTitle(message.data.data.title);
        }
      });
    }
  }, []);

  if (!title && !description) {
    return null;
  }

  return (
    <ModalBackground onClick={() => close()}>
      <ModalBody>
        <PageHeader style={{ fontSize: "4rem" }}>{title}</PageHeader>
        <div style={{ fontSize: "2rem", flex: 1, overflow: "auto" }}>
          {description}
        </div>
        <ActionButton>Close</ActionButton>
      </ModalBody>
    </ModalBackground>
  );
};

export default NotificationListener;
