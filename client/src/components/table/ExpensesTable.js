import { Link } from "react-router-dom";

function Datatable({ data }) {
  if (!Array.isArray(data)) {
    console.log(typeof data)
    if (typeof data === 'string') {
      return <div>{data}</div>;
    }
    return <div>Error: Data is not an array.</div>;
  }
  return (
    <div>

    <table>
      <thead>
        <tr>
          <th>ID</th>
          <th>Value</th>
          <th>Description</th>
          <th>Type</th>
          <th>Participants</th>
          <th>Status</th>
          <th>Recorrência</th>
        </tr>
      </thead>
      <tbody>
        {data.map((item) => (
          <tr key={item.ID}>
            <td>{item.ID}</td>
            <td>{item.value}</td>
            <td>{item.description}</td>
            <td>{item.type}</td>
            <td>
              {item.Participants && item.Participants.length > 0 ? (
                <ul className="no-bullets">
                  {item.Participants.map((participant) => (
                    <li key={participant.ID} style={{ listStyleType: 'none', alignContent: 'center' }}>
                      <Link to={`/participants/${participant.name}`}>
                        {participant.name}
                      </Link>
                    </li>
                  ))}
                </ul>
              ) : (
                "Sem participantes"
              )}
            </td>
            <td>{item.status === false ? "Não pago" : "Pago"}</td>
            <td>{item.is_active === false ? "Sim" : "Não"}</td>
          </tr>
        ))}
      </tbody>
    </table>
  </div>
  );
};

export default Datatable;
