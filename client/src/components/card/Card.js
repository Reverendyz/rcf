function Card({ title, description }) {
  return (
    <div className="card">
      <h1>{title}</h1>
      <p>{description}</p>
    </div>
  )
}

export default Card;
