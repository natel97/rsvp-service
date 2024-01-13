import { useParams } from "react-router-dom";
import { PageHeader } from "../../../components";
import { getAuthToken } from "../auth";
import { useEffect, useState } from "react";
import { GroupCard } from "../component";

const inviteGroup = (eventID, groupID) => {
  const apiKey = getAuthToken();
  fetch(`${import.meta.env.VITE_API_URL}/admin/invitation/group`, {
    method: "POST",
    headers: { Authorization: apiKey },
    body: JSON.stringify({
      eventID,
      groupID,
    }),
  });
};

const InviteGroup = () => {
  const apiKey = getAuthToken();
  const { id } = useParams();
  const [groups, setGroups] = useState([]);
  useEffect(() => {
    fetch(`${import.meta.env.VITE_API_URL}/admin/group`, {
      headers: { Authorization: apiKey },
    })
      .then((val) => val.json())
      .then((groups) => setGroups(groups));
  }, []);

  return (
    <div>
      <PageHeader>Invite to Event</PageHeader>
      {groups.map((group) => (
        <GroupCard
          key={group.ID}
          {...group}
          onClick={() => !group.InvitationID && inviteGroup(id, group.ID)}
        />
      ))}
    </div>
  );
};

export default InviteGroup;
