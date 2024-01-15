import { useEffect, useState } from "react";
import { getExistingInvitations } from "../utils/storeIDs";
import DelayedCard from "../components/delayedCard";
import { Card, PageHeader } from "../components";
import { useNavigate } from "react-router-dom";

const Home = () => {
  const [ids, setIDs] = useState([]);
  const navigate = useNavigate();

  useEffect(() => {
    const invitations = getExistingInvitations();
    setIDs(invitations);
  }, []);

  return (
    <div
      style={{ display: "flex", flexDirection: "column" }}
      className="full-height full-width"
    >
      <PageHeader>Your Events</PageHeader>
      <div style={{ overflow: "auto" }}>
        {ids.map((id) => (
          <DelayedCard
            onClick={() => navigate(`/invitation/${id}`)}
            id={id}
            key={id}
          />
        ))}
      </div>
      {!ids.length && <Card>No Invitations</Card>}
      <a
        target="_blank"
        href="https://github.com/natel97/rsvp-service"
        rel="noreferrer"
      >
        GitHub
      </a>
    </div>
  );
};

export default Home;
