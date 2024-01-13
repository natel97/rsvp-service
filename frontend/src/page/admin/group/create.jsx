import { useState } from "react";
import { ActionButton, PageHeader, TextInput } from "../../../components";
import { useNavigate } from "react-router-dom";
import { getAuthToken } from "../auth";

const CreateGroup = () => {
  const [Name, setName] = useState("");
  const apiKey = getAuthToken();
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
          value={Name}
          onChange={(e) => setName(e.target.value)}
        />
      </div>
      <ActionButton
        onClick={() => {
          fetch(`${import.meta.env.VITE_API_URL}/admin/group`, {
            method: "POST",
            body: JSON.stringify({
              Name,
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

export default CreateGroup;
