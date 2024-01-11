import { useParams } from "react-router-dom";
import { PageHeader } from "../../../components";
import { getAuthToken } from "../auth";
import { useEffect, useState } from "react";
import { PersonCard } from "../component";

const Invite = () => {
  const apiKey = getAuthToken();
  const { id } = useParams();
  const [people, setPeople] = useState([]);
  useEffect(() => {
    fetch(`${import.meta.env.VITE_API_URL}/admin/people`, {
      headers: { Authorization: apiKey },
    })
      .then((val) => val.json())
      .then((people) => setPeople(people));
  }, []);

  return (
    <div>
      <PageHeader>Invite to Event</PageHeader>
      {people.map((person) => (
        <PersonCard key={person.id} {...person} />
      ))}
    </div>
  );
};

export default Invite;
