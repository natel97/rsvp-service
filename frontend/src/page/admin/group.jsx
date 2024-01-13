import { useEffect, useState } from "react";
import { ActionButton, PageHeader } from "../../components";
import { getAuthToken } from "./auth";
import { GroupCard } from "./component";
import { useNavigate } from "react-router-dom";

const AdminGroups = () => {
  const apiKey = getAuthToken();
  const navigate = useNavigate();

  const [groups, setGroups] = useState([]);
  useEffect(() => {
    fetch(`${import.meta.env.VITE_API_URL}/admin/group`, {
      headers: {
        Authorization: apiKey,
      },
    }).then((val) => val.json().then((val) => setGroups(val)));
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
      <PageHeader>Manage Groups</PageHeader>
      <div style={{ flex: 1, overflow: "auto" }}>
        {groups.map((e) => (
          <GroupCard
            onClick={() => navigate(`./${e.ID}/people`)}
            key={e.id}
            {...e}
          />
        ))}
      </div>
      <div>
        <ActionButton onClick={() => navigate("./create")}>
          Create Group
        </ActionButton>
      </div>
    </div>
  );
};

export default AdminGroups;
