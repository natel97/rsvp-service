import { useEffect, useState } from "react";
import {
  ActionButton,
  EventCard,
  Modal,
  PageHeader,
  Text,
  TimeOption,
} from "../components";
import { useParams, useNavigate } from "react-router-dom";
import { storeInvitation } from "../utils/storeIDs";
import { NotifyButton } from "../components/notify";

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

const voteDay = (Acceptable, invitationID, optionID) => {
  fetch(
    `${
      import.meta.env.VITE_API_URL
    }/invitation/${invitationID}/time-selection/${optionID}`,
    {
      method: "PUT",
      body: JSON.stringify({
        Acceptable,
      }),
    }
  );
};

const SelectDays = ({ invitationID, timeOptions, open, setOpen }) => {
  return (
    <Modal open={open} setIsOpen={setOpen}>
      <PageHeader>Select Available Days</PageHeader>
      {timeOptions.map((option) => (
        <TimeOption
          option={option}
          key={option.id}
          onSelect={(acceptable) =>
            voteDay(acceptable, invitationID, option.id)
          }
        />
      ))}
      <ActionButton className="full-width" onClick={() => setOpen(false)}>
        Done
      </ActionButton>
    </Modal>
  );
};

const Invitation = () => {
  const params = useParams();
  const navigate = useNavigate();
  const [daySelectOpen, setDaySelectOpen] = useState(false);

  const [invitation, setInvitation] = useState(null);
  const needToSelectDays = invitation?.timeOptions.every(
    ({ isUpvote, isDownvote }) => !isUpvote && !isDownvote
  );

  const refreshData = () =>
    fetch(`${import.meta.env.VITE_API_URL}/invitation/${params.id}`).then(
      (val) => val.json().then((val) => setInvitation(val))
    );

  useEffect(() => {
    if (!params.id) {
      return;
    }

    refreshData();
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
      {invitation.invitationState === "PLANNING" && (
        <Text size="md">
          Select days you are available, then enable notifications to get
          notified when a day is decided.
        </Text>
      )}
      {invitation.subscribed && (
        <ActionButton
          onClick={() =>
            fetch(
              `${import.meta.env.VITE_API_URL}/invitation/${
                params.id
              }/subscribe`,
              { method: "DELETE" }
            )
          }
        >
          Stop Notifications 🔔
        </ActionButton>
      )}
      {!invitation.subscribed && (
        <NotifyButton
          url={`${import.meta.env.VITE_API_URL}/invitation/${
            params.id
          }/subscribe`}
        />
      )}
      {invitation.invitationState === "PLANNING" && (
        <ActionButton onClick={() => setDaySelectOpen(true)}>
          Select Available Dates 📅 {needToSelectDays && "❗❗"}
        </ActionButton>
      )}
      <SelectDays
        open={daySelectOpen}
        setOpen={(val) => setDaySelectOpen(val) || refreshData()}
        timeOptions={invitation.timeOptions}
        invitationID={params.id}
      />
      {invitation.date && (
        <ActionButton onClick={() => downloadEvent(params.id)}>
          Add to Calendar
        </ActionButton>
      )}
      {invitation?.invitationState !== "PLANNING" && (
        <ActionButton onClick={() => navigate("rsvp")}>
          <div>RSVP ({invitation.myAttendance || "No Response"})</div>
          <div>Friend ({invitation.myFriend || "No Response"})</div>
        </ActionButton>
      )}
    </div>
  );
};

export default Invitation;
