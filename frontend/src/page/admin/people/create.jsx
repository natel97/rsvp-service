import { useState } from "react";
import { ActionButton, PageHeader, TextInput } from "../../../components";
import { useNavigate } from "react-router-dom";
import { getAuthToken } from "../auth";

const CreatePerson = () => {
  const apiKey = getAuthToken();
  const [First, setFirst] = useState("");
  const [Last, setLast] = useState("");
  const navigate = useNavigate();

  return (
    <div>
      <PageHeader>Create Person</PageHeader>
      <div
        style={{
          display: "flex",
          flexDirection: "column",
          textAlign: "start",
          alignItems: "center",
        }}
      >
        <TextInput
          placeholder="First Name"
          value={First}
          onChange={(e) => setFirst(e.target.value)}
        />
        <TextInput
          placeholder="Last Name"
          value={Last}
          onChange={(e) => setLast(e.target.value)}
        />
      </div>
      <ActionButton
        onClick={() => {
          fetch(`${import.meta.env.VITE_API_URL}/admin/people`, {
            method: "POST",
            body: JSON.stringify({ First, Last }),
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

export default CreatePerson;
