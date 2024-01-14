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

const OptionEmoji = {
  Yes: "🟢",
  Maybe: "🤷‍♂️",
  No: "🔴",
  "": "🆕",
};

export const PersonCard = ({
  First,
  Last,
  InvitationID,
  InGroup,
  onClick,
  Going,
  BringingFriend,
}) => {
  return (
    <Card onClick={onClick} style={{ fontSize: "1.4rem" }}>
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
        <div style={{ display: "flex" }}>
          {InGroup && <div>[In Group]</div>}
          {(Going !== undefined || BringingFriend !== undefined) && (
            <div>
              {OptionEmoji[Going]}(+1{OptionEmoji[BringingFriend]})
            </div>
          )}
          {InvitationID && (
            <div
              onClick={(e) => {
                e.stopPropagation();
                navigator.clipboard.writeText(
                  `${window.location.origin}/invitation/${InvitationID}`
                );
              }}
            >
              (📋)
            </div>
          )}
        </div>
      </div>
    </Card>
  );
};

export const GroupCard = ({ Name, onClick }) => {
  return (
    <Card onClick={onClick}>
      <div
        style={{
          display: "flex",
          flexDirection: "row",
          justifyContent: "space-between",
        }}
      >
        <div>{Name}</div>
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
