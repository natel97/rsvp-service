import { useEffect, useState } from "react";
import { ActionButton, EventCard, PageHeader } from "../components";
import { useParams, useNavigate } from "react-router-dom";
import { storeInvitation } from "../utils/storeIDs";

const downloadEvent = (id) => {
  fetch(`${import.meta.env.VITE_API_URL}/invitation/${id}/download`).then(
    (val) => {
      const header = val.headers.get("Content-Disposition");
      const parts = header.split(";");
      const filename = parts[1].split("=")[1];
      val.text().then((data) => {
        var a = document.createElement("a");
        var url = window.URL.createObjectURL(
          new Blob([data], { type: "text/calendar" })
        );

        a.href = url;
        a.download = filename;
        document.body.append(a);
        a.click();
        window.URL.revokeObjectURL(url);
        a.remove();
      });
    }
  );
};

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

  storeInvitation(params.id);

  return (
    <div
      className="full-width full-height"
      style={{ display: "flex", flexDirection: "column" }}
    >
      <div className="full-width full-height">
        <PageHeader>You Are Invited</PageHeader>
        <EventCard {...invitation} />
      </div>
      <ActionButton onClick={() => downloadEvent(params.id)}>
        Add to Calendar
      </ActionButton>
      <ActionButton onClick={() => navigate("rsvp")}>
        <div>RSVP ({invitation.myAttendance || "No Response"})</div>
        <div>Friend ({invitation.myFriend || "No Response"})</div>
      </ActionButton>
    </div>
  );
};

export default Invitation;
