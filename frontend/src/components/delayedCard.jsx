import { useEffect, useState } from "react";
import { EventCard } from "../components";

const DelayedCard = ({ id = "", onClick }) => {
  const [invitation, setInvitation] = useState(null);

  useEffect(() => {
    if (!id) {
      return;
    }

    fetch(`${import.meta.env.VITE_API_URL}/invitation/${id}`).then((val) =>
      val.json().then((val) => setInvitation(val))
    );
  }, [id]);

  if (!invitation) {
    return <EventCard />;
  }

  return <EventCard onClick={onClick} {...invitation} />;
};

export default DelayedCard;
