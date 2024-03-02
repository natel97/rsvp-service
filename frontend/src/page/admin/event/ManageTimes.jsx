import { useState } from "react";
import {
  ActionButton,
  DAY_FORMAT,
  Modal,
  PageHeader,
  TIME_FORMAT,
  Text,
} from "../../../components";
import moment from "moment";
import { getAuthToken } from "../auth";

const updateTime = (eventID, timeID) => {
  const apiKey = getAuthToken();
  fetch(
    `${
      import.meta.env.VITE_API_URL
    }/admin/event/${eventID}/select-time/${timeID}`,
    {
      method: "PUT",
      headers: {
        Authorization: apiKey,
      },
    }
  );
};

const deleteOption = (optionID) => {
  const apiKey = getAuthToken();
  fetch(`${import.meta.env.VITE_API_URL}/admin/event/time-option/${optionID}`, {
    method: "DELETE",
    headers: {
      Authorization: apiKey,
    },
  });
};

const parseTime = (time) =>
  moment(time).format(DAY_FORMAT) + " " + moment(time).format(TIME_FORMAT);

const ManageTimes = ({ ID, TimeOptions = [] }) => {
  const [open, setOpen] = useState(false);

  return (
    <>
      <ActionButton onClick={() => setOpen(true)}>Manage Times</ActionButton>
      <Modal open={open} setIsOpen={setOpen}>
        <PageHeader>Manage Times</PageHeader>
        {TimeOptions.map((option) => (
          <div key={option.id}>
            <Text size="md">{parseTime(option.time)}</Text>
            <div style={{ display: "flex", alignItems: "center" }}>
              <Text size="md">👍️{option.upvote}</Text>
              <Text size="md">👎️{option.downvote}</Text>
              <ActionButton
                style={{ marginLeft: "auto" }}
                onClick={() => updateTime(ID, option.id)}
              >
                Select
              </ActionButton>
              <ActionButton onClick={() => deleteOption(option.id)}>
                🚮️
              </ActionButton>
            </div>
          </div>
        ))}
        <ActionButton className="full-width" onClick={() => setOpen(false)}>
          Close
        </ActionButton>
      </Modal>
    </>
  );
};

export default ManageTimes;
