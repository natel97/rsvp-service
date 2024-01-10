import { useEffect, useState } from "react";
import { ActionButton, EventCard, Option, PageHeader } from "../components";
import { useNavigate, useParams } from "react-router-dom";

const submit = async (id, going, bringingAFriend) => {
  return fetch(`${import.meta.env.VITE_API_URL}/invitation/${id}/rsvp`, {
    method: "POST",
    body: JSON.stringify({ going, bringingAFriend }),
  });
};

const RSVP = () => {
  const navigate = useNavigate();
  const params = useParams();
  const [invitation, setInvitation] = useState(null);
  const [going, setGoing] = useState();
  const [friend, setFriend] = useState();

  useEffect(() => {
    if (invitation === null) {
      return;
    }

    setGoing(invitation.myAttendance);
    setFriend(invitation.myFriend);
  }, [invitation]);

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
      className="full-height full-width"
      style={{
        display: "flex",
        flexDirection: "column",
        justifyContent: "space-between",
      }}
    >
      <div style={{ overflow: "auto" }}>
        <PageHeader>RSVP</PageHeader>
        <EventCard {...invitation} attendance={null} />
        <h2>Are You attending?</h2>
        <Option
          prefix="going"
          options={["Yes", "Maybe", "No"]}
          current={going}
          setCurrent={setGoing}
        />
        <br />
        <h2>Are You Bringing a Friend?</h2>
        <Option
          prefix="friend"
          options={["Yes", "Maybe", "No"]}
          current={friend}
          setCurrent={setFriend}
        />
      </div>
      <ActionButton
        onClick={() => {
          if (
            going === invitation.myAttendance &&
            friend === invitation.myFriend
          ) {
            navigate("./..");
            return;
          }
          submit(params.id, going, friend).then(() => navigate("./.."));
        }}
        style={{ marginBottom: "0" }}
      >
        Submit
      </ActionButton>
    </div>
  );
};

export default RSVP;
