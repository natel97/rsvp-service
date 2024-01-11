import { useState } from "react";
import { ActionButton, PageHeader, TextInput } from "../../../components";
import { getAuthToken } from "../auth";
import { useNavigate } from "react-router-dom";
import moment from "moment";

const CreateEvent = () => {
  const apiKey = getAuthToken();
  const [City, setCity] = useState("");
  const [date, setDate] = useState("");
  const [Description, setDescription] = useState("");
  const [InternalNote, setInternalNote] = useState("");
  const [Street, setStreet] = useState("");
  const [Title, setTitle] = useState("");
  const navigate = useNavigate();

  return (
    <div>
      <PageHeader>Create Event</PageHeader>
      <div
        style={{
          display: "flex",
          flexDirection: "column",
          textAlign: "start",
          alignItems: "center",
        }}
      >
        <TextInput
          placeholder="Title"
          value={Title}
          onChange={(e) => setTitle(e.target.value)}
        />
        <TextInput
          placeholder="City, State Postal Code"
          value={City}
          onChange={(e) => setCity(e.target.value)}
        />
        <TextInput
          placeholder="Unit, Number, Street"
          value={Street}
          onChange={(e) => setStreet(e.target.value)}
        />
        <TextInput
          type="datetime-local"
          placeholder="Date"
          value={date}
          min={new Date()}
          onChange={(e) => setDate(e.target.value)}
        />
        <TextInput
          placeholder="Public Description"
          value={Description}
          onChange={(e) => setDescription(e.target.value)}
        />
        <TextInput
          placeholder="Internal Description"
          value={InternalNote}
          onChange={(e) => setInternalNote(e.target.value)}
        />
      </div>
      <ActionButton
        onClick={() => {
          fetch(`${import.meta.env.VITE_API_URL}/admin/event`, {
            method: "POST",
            body: JSON.stringify({
              City,
              Date: moment(date).toDate(),
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
