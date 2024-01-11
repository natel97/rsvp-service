import styled from "styled-components";
import { DAY_FORMAT, TIME_FORMAT } from "../../components";
import moment from "moment";

const Card = styled.div`
  background: #111;
  padding: 8px;
  border-radius: 8px;
  margin: 12px 32px;
  font-size: 1.5rem;
  cursor: pointer;
`;

export const PersonCard = ({ First, Last, InvitationID, onClick }) => {
  return (
    <Card onClick={onClick}>
      <div
        style={{
          display: "flex",
          flexDirection: "row",
          justifyContent: "space-between",
        }}
      >
        <div>
          {First} {Last}
        </div>
        {InvitationID && (
          <div
            onClick={(e) => {
              e.stopPropagation();
              navigator.clipboard.writeText(
                `${window.location.origin}/invitation/${InvitationID}`
              );
            }}
          >
            [Copy]
          </div>
        )}
      </div>
    </Card>
  );
};

const TimeStamp = ({ time }) => {
  const dayString = moment(time).format(DAY_FORMAT);
  const timeString = moment(time).format(TIME_FORMAT);

  return (
    <div>
      <div>{dayString}</div>
      <div>{timeString}</div>
    </div>
  );
};

export const EventCard = ({
  City,
  Date,
  Description,
  InternalNote,
  Street,
  Title,
  onClick,
}) => {
  return (
    <Card onClick={onClick}>
      <div>{Title}</div>
      <div>
        {Street}, {City}
      </div>
      <TimeStamp time={Date} />
      <div>{Description}</div>
      {InternalNote && <div>Internal Note - Not Shared</div>}
      <div>{InternalNote}</div>
    </Card>
  );
};
