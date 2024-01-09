import { useState } from "react";
import { Card, EventCard, Option, PageHeader } from "../components";

const RSVP = () => {
  const [going, setGoing] = useState();
  const [friend, setFriend] = useState();
  return (
    <div className="full-height full-width">
      <PageHeader>RSVP</PageHeader>
      <EventCard attendance={null} />
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
      <Card>
        <h2 style={{ textAlign: "center" }}>Submit</h2>
      </Card>
    </div>
  );
};

export default RSVP;
