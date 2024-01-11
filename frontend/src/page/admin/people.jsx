import { useEffect, useState } from "react";
import { ActionButton, PageHeader } from "../../components";
import { getAuthToken } from "./auth";
import { PersonCard } from "./component";
import { useNavigate } from "react-router-dom";

const AdminPeople = () => {
  const apiKey = getAuthToken();
  const navigate = useNavigate();

  const [people, setPeople] = useState([]);
  useEffect(() => {
    fetch(`${import.meta.env.VITE_API_URL}/admin/people`, {
      headers: {
        Authorization: apiKey,
      },
    }).then((val) => val.json().then((val) => setPeople(val)));
  }, [apiKey]);

  return (
    <div
      className="full-height full-width"
      style={{
        display: "flex",
        flexDirection: "column",
        justifyContent: "space-between",
      }}
    >
      <PageHeader>Manage People</PageHeader>
      <div style={{ flex: 1, overflow: "auto" }}>
        {people.map((e) => (
          <PersonCard key={e.id} {...e} />
        ))}
      </div>
      <div>
        <ActionButton onClick={() => navigate("./create")}>
          Create Person
        </ActionButton>
      </div>
    </div>
  );
};

export default AdminPeople;
