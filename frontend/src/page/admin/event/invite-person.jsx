import { useParams } from "react-router-dom";
import { PageHeader } from "../../../components";
import { getAuthToken } from "../auth";
import { useEffect, useState } from "react";
import { PersonCard } from "../component";

const invitePerson = (eventID, personID) => {
  const apiKey = getAuthToken();
  fetch(`${import.meta.env.VITE_API_URL}/admin/invitation`, {
    method: "POST",
    headers: { Authorization: apiKey },
    body: JSON.stringify({
      eventID,
      personID,
    }),
  });
};

const InvitePerson = () => {
  const apiKey = getAuthToken();
  const { id } = useParams();
  const [people, setPeople] = useState([]);
  useEffect(() => {
    fetch(`${import.meta.env.VITE_API_URL}/admin/event/${id}/people`, {
      headers: { Authorization: apiKey },
    })
      .then((val) => val.json())
      .then((people) => setPeople(people));
  }, []);

  return (
    <div>
      <PageHeader>Invite to Event</PageHeader>
      {people.map((person) => (
        <PersonCard
          key={person.ID}
          {...person}
          onClick={() => !person.InvitationID && invitePerson(id, person.ID)}
        />
      ))}
    </div>
  );
};

export default InvitePerson;
