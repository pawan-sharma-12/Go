package greetings
import( 
    "fmt"
    "errors"
    "math/rand"
) 
func Hello(name string) (string, error) {
    //error if no name 
    if name == "" {
        return "", errors.New("empty name")
    }
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}
func Hellos(names []string)(map[string]string, error){
    messages := make(map[string]string)
    for _, name := range names {
        message, err := Hello(name)
        if err != nil {
            return nil, err
        }
        messages[name] = message

    }
    return messages , nil
}
func randomFormat() string {
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
        "Hey, %v! Welcome buddy!",
        "Yup i got you , %v! Welcome here!",
        "Yo buddy, %v! whats poppin!",
        "Haha chill, %v! Welcome onboard dude!",
    }
    return formats[rand.Intn(len(formats))]
}