import { useState } from "react";
import { ActionButton, PageHeader, TextInput } from "../../../components";
import { getAuthToken } from "../auth";
import { useNavigate } from "react-router-dom";

export const getDefaultDateTime = () => {
  const fullDate = new Date().toISOString();
  const date = fullDate.split("T")[0];

  return `${date}T18:30:00.000`;
};

const CreateEvent = () => {
  const apiKey = getAuthToken();
  const [City, setCity] = useState("");
  const [Description, setDescription] = useState("");
  const [InternalNote, setInternalNote] = useState("");
  const [Street, setStreet] = useState("");
  const [Title, setTitle] = useState("");
  const navigate = useNavigate();

  return (
    <div
      className="full-height full-width"
      style={{ display: "flex", flexDirection: "column" }}
    >
      <PageHeader>Create Event</PageHeader>
      <div
        style={{
          display: "flex",
          flexDirection: "column",
          textAlign: "start",
          alignItems: "stretch",
          flex: 1,
        }}
      >
        <TextInput
          placeholder="Title"
          value={Title}
          onChange={(e) => setTitle(e.target.value)}
        />
        <TextInput
          placeholder="Unit, Number, Street"
          value={Street}
          onChange={(e) => setStreet(e.target.value)}
        />
        <TextInput
          placeholder="City, State Postal Code"
          value={City}
          onChange={(e) => setCity(e.target.value)}
        />
        <TextInput
          placeholder="Public Description"
          value={Description}
          area
          onChange={(e) => setDescription(e.target.value)}
        />
        <TextInput
          placeholder="Internal Description"
          area
          value={InternalNote}
          onChange={(e) => setInternalNote(e.target.value)}
        />
      </div>
      <ActionButton
        className="full-width"
        onClick={() => {
          fetch(`${import.meta.env.VITE_API_URL}/admin/event`, {
            method: "POST",
            body: JSON.stringify({
              City,
              Description,
              InternalNote,
              Street,
              Title,
            }),
            headers: {
              Authorization: apiKey,
            },
          }).then(() => navigate("./.."));
        }}
      >
        Create
      </ActionButton>
    </div>
  );
};

export default CreateEvent;
