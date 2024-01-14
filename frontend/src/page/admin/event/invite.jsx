import { useNavigate, useParams } from "react-router-dom";
import { ActionButton, EventCard, PageHeader } from "../../../components";
import { PersonCard } from "../component";
import { useEffect, useState } from "react";
import { getAuthToken } from "../auth";

const Invite = () => {
  const navigate = useNavigate();
  const id = useParams().id;
  const apiKey = getAuthToken();
  const [attendance, setAttendance] = useState({});
  useEffect(() => {
    fetch(`${import.meta.env.VITE_API_URL}/admin/event/${id}/attendance`, {
      headers: { Authorization: apiKey },
    })
      .then((val) => val.json())
      .then((val) => setAttendance(val));
  }, []);

  return (
    <div
      className="full-height"
      style={{ display: "flex", flexDirection: "column" }}
    >
      <PageHeader>Manage Event</PageHeader>
      <EventCard
        city={attendance.City}
        date={attendance.Date}
        street={attendance.Street}
        description={`${attendance.Description}\n${
          attendance.InternalNote ? "Internal \n" + attendance.InternalNote : ""
        }`}
        title={attendance.Title}
        attendance={null}
      />
      <div style={{ flex: 1, overflow: "auto" }}>
        {attendance.Attendance?.map((person) => (
          <PersonCard key={person.ID} {...person} />
        ))}
      </div>
      <div>
        <ActionButton onClick={() => navigate("./person")}>
          Invite Person
        </ActionButton>
        <ActionButton onClick={() => navigate("./group")}>
          Invite Group
        </ActionButton>
      </div>
    </div>
  );
};

export default Invite;
