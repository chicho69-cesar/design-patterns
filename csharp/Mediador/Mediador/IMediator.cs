namespace Mediador {
    public interface IMediator {
        public void Send(string message, Colleague colleague);
    }
}