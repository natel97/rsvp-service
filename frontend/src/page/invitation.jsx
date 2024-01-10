import { useEffect, useState } from "react";
import { ActionButton, EventCard, PageHeader } from "../components";
import { useParams, useNavigate } from "react-router-dom";

const Invitation = () => {
  const params = useParams();
  const navigate = useNavigate();
  const [invitation, setInvitation] = useState(null);

  useEffect(() => {
    if (!params.id) {
      return;
    }

    fetch(`${import.meta.env.VITE_API_URL}/invitation/${params.id}`).then(
      (val) => val.json().then((val) => setInvitation(val))
    );
  }, [params.id]);

  if (invitation === null) {
    return <div>loading</div>;
  }

  return (
    <div
      className="full-width full-height"
      style={{ display: "flex", flexDirection: "column" }}
    >
      <div className="full-width full-height">
        <PageHeader>You Are Invited</PageHeader>
        <EventCard {...invitation} />
      </div>
      <ActionButton>Add to Calendar</ActionButton>
      <ActionButton onClick={() => navigate("rsvp")}>
        <div>RSVP ({invitation.myAttendance || "No Response"})</div>
        <div>Friend ({invitation.myFriend || "No Response"})</div>
      </ActionButton>
    </div>
  );
};

export default Invitation;
