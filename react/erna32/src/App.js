import './App.css';
import './styles.css';
import Card from './Card';
import Song from './elan_piano.mp3';

function App() {
  function play(){
    new Audio(Song).play();
  }
  return (
    <div className="App" onLoad={play}>
      

     <Card onClick={play} />
     <audio className="audioController" src={Song} controls autoplay />
    </div>
  );
}

export default App;
