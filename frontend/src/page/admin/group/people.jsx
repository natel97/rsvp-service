import { useParams } from "react-router-dom";
import { getAuthToken } from "../auth";
import { useEffect, useState } from "react";
import { PageHeader } from "../../../components";
import { PersonCard } from "../component";

const editPersonGroup = (groupID, personID, inGroup = false) => {
  const apiKey = getAuthToken();
  fetch(
    `${import.meta.env.VITE_API_URL}/admin/group/${groupID}/person/${personID}`,
    {
      method: inGroup ? "DELETE" : "POST",
      headers: { Authorization: apiKey },
    }
  );
};

const ManageGroupPeople = () => {
  const apiKey = getAuthToken();
  const { id } = useParams();
  const [people, setPeople] = useState([]);
  useEffect(() => {
    fetch(`${import.meta.env.VITE_API_URL}/admin/group/${id}/person`, {
      headers: { Authorization: apiKey },
    })
      .then((val) => val.json())
      .then((people) => setPeople(people));
  }, []);

  return (
    <div>
      <PageHeader>Manage Group</PageHeader>
      {people.map((person) => (
        <PersonCard
          key={person.ID}
          {...person}
          onClick={() => editPersonGroup(id, person.ID, person.InGroup)}
        />
      ))}
    </div>
  );
};

export default ManageGroupPeople;
