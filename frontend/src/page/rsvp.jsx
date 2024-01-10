import { useEffect, useState } from "react";
import { ActionButton, EventCard, Option, PageHeader } from "../components";
import { useNavigate, useParams } from "react-router-dom";

const RSVP = () => {
  const navigate = useNavigate();
  const params = useParams();
  const [invitation, setInvitation] = useState(null);
  const [going, setGoing] = useState();
  const [friend, setFriend] = useState();

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
      <ActionButton style={{ marginBottom: "0" }}>Submit</ActionButton>
    </div>
  );
};

export default RSVP;
