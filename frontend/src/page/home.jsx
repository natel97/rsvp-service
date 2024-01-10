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
    <div className="full-height full-width">
      <PageHeader>Your Events</PageHeader>
      {ids.map((id) => (
        <DelayedCard
          onClick={() => navigate(`/invitation/${id}`)}
          id={id}
          key={id}
        />
      ))}
      {!ids.length && <Card>No Invitations</Card>}
    </div>
  );
};

export default Home;
