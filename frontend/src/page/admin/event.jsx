import { useEffect, useState } from "react";
import { ActionButton, PageHeader } from "../../components";
import { getAuthToken } from "./auth";
import { EventCard } from "./component";
import { useNavigate } from "react-router-dom";

const AdminEvent = () => {
  const apiKey = getAuthToken();
  const navigate = useNavigate();

  const [events, setEvents] = useState([]);
  useEffect(() => {
    fetch(`${import.meta.env.VITE_API_URL}/admin/event`, {
      headers: {
        Authorization: apiKey,
      },
    }).then((val) => val.json().then((val) => setEvents(val)));
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
      <PageHeader>Manage Events</PageHeader>
      <div style={{ flex: 1, overflow: "auto" }}>
        {events.map((e) => (
          <EventCard
            onClick={() => navigate(`./${e.ID}/invite`)}
            key={e.id}
            {...e}
          />
        ))}
      </div>
      <div>
        <ActionButton onClick={() => navigate("./create")}>
          Create Event
        </ActionButton>
      </div>
    </div>
  );
};

export default AdminEvent;
