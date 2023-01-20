namespace Strategy {
    public class StrategyContext {
        private IAlcoholic alcoholic;

        public enum Behaviour {
            Eyes,
            Invite,
            Prince
        }

        public StrategyContext() {
            this.alcoholic = new StrategyEyes();
        }

        public void SetEyes() {
            this.alcoholic = new StrategyEyes();
        }

        public void SetInvite() {
            this.alcoholic = new StrategyInvite();
        }

        public void SetPrince() {
            this.alcoholic = new StrategyBePrince();
        }

        public void BeRomantic() {
            this.alcoholic.BeRomantic();
        }

        public void BeRomantic(Behaviour option) {
            switch (option) {
                case Behaviour.Eyes:
                    this.alcoholic = new StrategyEyes();
                    break;
                case Behaviour.Invite:
                    this.alcoholic = new StrategyInvite();
                    break;
                case Behaviour.Prince:
                    this.alcoholic = new StrategyBePrince();
                    break;
            }

            this.alcoholic.BeRomantic();
        }
    }
}