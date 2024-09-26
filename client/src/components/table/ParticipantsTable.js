import { FaPlus } from "react-icons/fa";
import { Link } from "react-router-dom";

function ParticipantsTable({ data }) {
  if (!Array.isArray(data)) {
    console.log(typeof data)
    if (typeof data === 'string') {
      return <div>{data}</div>;
    }
    return <div>Error: Data is not an array.</div>;
  }
  return (
    <div className="table-container">
      <table>
        <thead>
          <tr>
            <th>Name</th>
            <th>Contas</th>
            <th>Ativo</th>
          </tr>
        </thead>
        <tbody>
          {data.map((item) => (
            <tr key={item.ID}>
              <td>{item.name}</td>
              <td>
                {item.Expenses && item.Expenses.length > 0 ? (
                  <ul className="no-bullets">
                    {item.Expenses.map((expense) => (
                      <li key={expense.ID} style={{ listStyleType: 'none', alignContent: 'center' }}>
                        <Link to={`/expenses/${expense.ID}`}>
                          {expense.description !== "" ? expense.description : expense.ID}
                        </Link>
                      </li>
                    ))}
                  </ul>
                ) : (
                  "Sem Contas"
                )}
              </td>
              <td>{item.is_active === false ? "Sim" : "Não"}</td>
            </tr>
          ))}
        </tbody>
      </table>
      <button className="add-button">
        <FaPlus /> Adicionar Participante
      </button>
    </div>
  );
};

export default ParticipantsTable;
