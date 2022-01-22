from playsound import playsound
from pedalboard import (
    Pedalboard,
    Convolution,
    Compressor,
    Distortion,
    Chorus,
    Gain,
    Reverb,
    Limiter,
    Phaser,
)
# import os
# import requests
# import json
import soundfile as sf

def main():
    try:
        audio_file, sample_rate = sf.read('speech_sample.wav')
        board_phaser_reverb = Pedalboard([
            Phaser(),
            Reverb(),
        ], sample_rate=sample_rate)
        audio_with_added_effect_phaser_reverb = board_phaser_reverb(audio_file)
        audio_with_added_effect_phaser_reverb_file_name = 'speech_sample_effect_phaser_reverb.wav'
        with sf.SoundFile(audio_with_added_effect_phaser_reverb_file_name, 'w', samplerate=sample_rate, channels=len(audio_with_added_effect_phaser_reverb.shape)) as f:
            f.write(audio_with_added_effect_phaser_reverb)

        board_distortion_chorus_gain = Pedalboard([
            Distortion(),
            Chorus(),
            Gain(),
        ], sample_rate=sample_rate)
        audio_with_added_effect_distortion_chorus_gain = board_distortion_chorus_gain(audio_file)
        audio_with_added_effect_distortion_chorus_gain_file_name = 'speech_sample_effect_distortion_chorus_gain.wav'
        with sf.SoundFile(audio_with_added_effect_distortion_chorus_gain_file_name, 'w', samplerate=sample_rate,
                          channels=len(audio_with_added_effect_distortion_chorus_gain.shape)) as f:
            f.write(audio_with_added_effect_distortion_chorus_gain)

        # playsound(audio_with_added_effect)
    except Exception as e:
        print(e)

# Press the green button in the gutter to run the script.
if __name__ == '__main__':
    main()